const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function putFileServices() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4410";
  const accountName = "sto8607";
  const parameters = {
    cors: {
      corsRules: [
        {
          allowedHeaders: ["x-ms-meta-abc", "x-ms-meta-data*", "x-ms-meta-target*"],
          allowedMethods: ["GET", "HEAD", "POST", "OPTIONS", "MERGE", "PUT"],
          allowedOrigins: ["http://www.contoso.com", "http://www.fabrikam.com"],
          exposedHeaders: ["x-ms-meta-*"],
          maxAgeInSeconds: 100,
        },
        {
          allowedHeaders: ["*"],
          allowedMethods: ["GET"],
          allowedOrigins: ["*"],
          exposedHeaders: ["*"],
          maxAgeInSeconds: 2,
        },
        {
          allowedHeaders: ["x-ms-meta-12345675754564*"],
          allowedMethods: ["GET", "PUT"],
          allowedOrigins: ["http://www.abc23.com", "https://www.fabrikam.com/*"],
          exposedHeaders: ["x-ms-meta-abc", "x-ms-meta-data*", "x-ms-meta-target*"],
          maxAgeInSeconds: 2000,
        },
      ],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileServices.setServiceProperties(
    resourceGroupName,
    accountName,
    parameters
  );
  console.log(result);
}

putFileServices().catch(console.error);
