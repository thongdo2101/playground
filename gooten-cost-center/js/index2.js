const axios = require("axios").default;
const fs = require("fs");
const { v4: uuidv4 } = require('uuid')

async function getAllSku() {
  try {
    let url =
      "https://api.print.io/api/v/5/source/api/products/?recipeid=5ccd6548-ba59-4a13-98cc-a4970ba15d87&languageCode=en&countryCode=US&all=true";
    let res = await axios.get(url);
    let vendorProducts = new Array()
    let productTypes = new Array()
    let costCenters = new Array()
    for (let i = 0; i < res.data.Products.length; i++) {
      let product = res.data.Products[i];
      let id = product.Id;
      let name = product.Name;
      let fileName = name
        .replace(" ", "_")
        .replace(/[\W_]+/g, "_")
        .replace(/^_+|_+$/g, "")
        .toLowerCase();

      let productUrl = `https://gtnadminassets.blob.core.windows.net/productdatav3/${fileName}.json`;
      let skusUrl = `https://gtnadminassets.blob.core.windows.net/productdatav3/${fileName}_skus.json`;
      let shippingUrl = `https://gtnadminassets.blob.core.windows.net/productpricing/${fileName}_shipping.json`;
      //
      let productRes
      try {
        productRes = await axios.get(productUrl);
      } catch {}
      let options = new Array()
      for (let j = 0; j < productRes.data.product.regions.length; j++) {
        let region = productRes.data.product.regions[j];
        for (let m = 0; m < region["sub-categories"].length; m++) {
          let subCate = region["sub-categories"][m];
          for (let n = 0; n < subCate.attributes.length; n++) {
            let attr = subCate.attributes[n];
            if (attr.id != undefined) {
              options.push({
                "id": attr.id,
                "name": attr.name,
                "values": attr.values
              })
            }
          }
        }
      }
      //
      let skusRes;
      try {
        skusRes = await axios.get(skusUrl);
      } catch {}
      let product_types = new Array();
      if (skusRes != undefined) {
        for (let j = 0; j < skusRes.data.length; j++) {
          let skuItem = skusRes.data[j];
          let sku = skuItem.sku;
          let minCost = skuItem.minCost;
          let maxCost = skuItem.maxCost;
          let skuOptions = new Array()
          //
          let keys = Object.keys(skuItem)
          for (let k = 0; k < keys.length; k++) {
            let key = keys[k];
            let value = skuItem[key]
            options.forEach(option => {
              if (option.id == key) {
                let optionName = option.name
                option.values.forEach(val => {
                  if (val.id == value) {
                    let valueName = val.name
                    skuOptions.push({
                      "option_title": optionName,
                      "option_value": valueName
                    })
                  }
                })
              }
            })
          }
          product_types.push({
            name: sku,
            code: sku,
            sku: sku,
            minCost,
            maxCost,
            "custom_options": skuOptions
          });
        }
      }
      // shipping fee
      let shippingFeeRes;
      try {
        shippingFeeRes = await axios.get(shippingUrl);
      } catch {}
      let metadata 
      if (shippingFeeRes != undefined) {
        metadata = {
          shipping_fee:
            shippingFeeRes != undefined && shippingFeeRes.data != undefined
              ? shippingFeeRes.data
              : null,
        };
      }
      let vendorProduct = {
        id: uuidv4(),
        name,
        vendor: "Gooten",
        productId: id,
        metadata
      }
      vendorProducts.push(vendorProduct)
      product_types.forEach(pType => {
        let productType = {
          id: uuidv4(),
          vendorProductID: vendorProduct.id,
          name: pType.name,
          code: pType.code,
          sku: pType.sku
        }
        productTypes.push(productType)
        //
        let costCenter = {
          id: uuidv4(),
          vendor_product_id: vendorProduct.id,
          sku: pType.sku,
          base_cost: pType.minCost,
          custom_options: pType.custom_options.map(option => {
            option.store_option_title = option.option_title
            option.store_option_value = option.option_value
            return option
          })
        }
        costCenters.push(costCenter)
      })
      vendorProduct.costCenters = costCenters
      console.log(id)
    }
    // call api
    
    // write files
    // fs.writeFile(`./data/final/vendor_products.json`, JSON.stringify(vendorProducts, null, 2), (err) => {
    //   if (err) {
    //   }
    // });
    // //
    // fs.writeFile(`./data/final/product_types.json`, JSON.stringify(productTypes, null, 2), (err) => {
    //   if (err) {
    //   }
    // });
    // //
    // fs.writeFile(`./data/final/cost_centers.json`, JSON.stringify(costCenters, null, 2), (err) => {
    //   if (err) {
    //   }
    // });
    //
    console.log("finish");
  } catch (err) {
    console.log(err);
  }
}

getAllSku();
