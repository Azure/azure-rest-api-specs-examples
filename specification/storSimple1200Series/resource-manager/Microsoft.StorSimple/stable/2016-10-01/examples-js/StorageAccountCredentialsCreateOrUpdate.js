const { StorSimpleManagementClient } = require("@azure/arm-storsimple1200series");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the storage account credential
 *
 * @summary Creates or updates the storage account credential
 * x-ms-original-file: specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/StorageAccountCredentialsCreateOrUpdate.json
 */
async function storageAccountCredentialsCreateOrUpdate() {
  const subscriptionId = "9eb689cd-7243-43b4-b6f6-5c65cb296641";
  const credentialName = "DummySacForSDKTest";
  const resourceGroupName = "ResourceGroupForSDKTest";
  const managerName = "hAzureSDKOperations";
  const storageAccount = {
    name: "DummySacForSDKTest",
    accessKey: {
      encryptionAlgorithm: "RSAES_PKCS1_v_1_5",
      encryptionCertificateThumbprint: "D73DB57C4CDD6761E159F8D1E8A7D759424983FD",
      value:
        "Ev1tm0QBmpGGm4a58GkqLqx8veJEEgQtg5K3Jizpmy7JdSv9dlcRwk59THw6KIdMDlEHcS8mPyneBtOEQsh4wkcFB7qrmQz+KsRAyIhEm6bwPEm3qN8+aDDzNcXn/6vu/sqV0AP7zit9/s7SxXGxjKrz4zKnOy16/DbzRRmUHNO+HO6JUM0cUfHXTX0mEecbsXqBq0A8IEG8z+bJgXX1EhoGkzE6yVsObm4S1AcKrLiwWjqmSLji5Q8gGO+y4KTTmC3p45h5GHHXjJyOccHhySWDAffxnTzUD/sOoh+aD2VkAYrL3DdnkVzhAdfcZfVI4soONx7tYMloZIVsfW1M2Q==",
    },
    cloudType: "Azure",
    enableSSL: "Enabled",
    endPoint: "blob.core.windows.net",
    location: "West US",
    login: "SacForSDKTest",
  };
  const credential = new DefaultAzureCredential();
  const client = new StorSimpleManagementClient(credential, subscriptionId);
  const result = await client.storageAccountCredentials.beginCreateOrUpdateAndWait(
    credentialName,
    resourceGroupName,
    managerName,
    storageAccount
  );
  console.log(result);
}

storageAccountCredentialsCreateOrUpdate().catch(console.error);
