const { InformaticaDataManagement } = require("@azure/arm-informaticadatamanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a InformaticaServerlessRuntimeResource
 *
 * @summary Update a InformaticaServerlessRuntimeResource
 * x-ms-original-file: specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/ServerlessRuntimes_Update_MaximumSet_Gen.json
 */
async function serverlessRuntimesUpdate() {
  const subscriptionId =
    process.env["INFORMATICA_SUBSCRIPTION_ID"] || "3599DA28-E346-4D9F-811E-189C0445F0FE";
  const resourceGroupName = process.env["INFORMATICA_RESOURCE_GROUP"] || "rgopenapi";
  const organizationName = "W5";
  const serverlessRuntimeName = "t_";
  const properties = {
    properties: {
      description: "ocprslpljoikxyduackzqnkuhyzrh",
      advancedCustomProperties: [{ key: "qcmc", value: "unraxmnohdmvutt" }],
      applicationType: "CDI",
      computeUnits: "uncwbpu",
      executionTimeout: "tjyfytuywriabt",
      platform: "AZURE",
      serverlessAccountLocation: "goaugkyfanqfnvcmntreibqrswfpis",
      serverlessRuntimeConfig: {
        cdiConfigProps: [
          {
            applicationConfigs: [
              {
                name: "upfvjrqcrwwedfujkmsodeinw",
                type: "lw",
                customized: "j",
                defaultValue: "zvgkqwmi",
                platform: "dixfyeobngivyvf",
                value: "mozgsetpwjmtyl",
              },
            ],
            engineName: "hngsdqvtjdhwqlbqfotipaiwjuys",
            engineVersion: "zlrlbg",
          },
        ],
        cdieConfigProps: [
          {
            applicationConfigs: [
              {
                name: "upfvjrqcrwwedfujkmsodeinw",
                type: "lw",
                customized: "j",
                defaultValue: "zvgkqwmi",
                platform: "dixfyeobngivyvf",
                value: "mozgsetpwjmtyl",
              },
            ],
            engineName: "hngsdqvtjdhwqlbqfotipaiwjuys",
            engineVersion: "zlrlbg",
          },
        ],
      },
      serverlessRuntimeNetworkProfile: {
        networkInterfaceConfiguration: {
          subnetId:
            "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/subnet1",
          vnetId:
            "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/HypernetVnet1",
          vnetResourceGuid: "5328d299-1462-4be0-bef1-303a28e556a0",
        },
      },
      serverlessRuntimeTags: [{ name: "korveuycuwhs", value: "uyiuegxnkgp" }],
      serverlessRuntimeUserContextProperties: {
        userContextToken: "ctgebtvjhwh",
      },
      supplementaryFileLocation: "csxaqzpxu",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new InformaticaDataManagement(credential, subscriptionId);
  const result = await client.serverlessRuntimes.update(
    resourceGroupName,
    organizationName,
    serverlessRuntimeName,
    properties,
  );
  console.log(result);
}
