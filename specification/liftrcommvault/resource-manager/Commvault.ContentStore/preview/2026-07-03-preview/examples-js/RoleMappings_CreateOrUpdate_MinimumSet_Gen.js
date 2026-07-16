const { ContentStoreClient } = require("@azure/arm-commvaultcontentstore");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a RoleMapping
 *
 * @summary create a RoleMapping
 * x-ms-original-file: 2026-07-03-preview/RoleMappings_CreateOrUpdate_MinimumSet_Gen.json
 */
async function roleMappingsCreateOrUpdateMinimumSetSingleBackupAdminRoleOnly() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "1B4766AF-8D4B-4B44-9CF1-587E003DFF7F";
  const client = new ContentStoreClient(credential, subscriptionId);
  const result = await client.roleMappings.createOrUpdate("rgcommvault", "myCloudAccount", {
    properties: {
      roles: [
        {
          roleName: "BackupAdmin",
          entities: [
            {
              id: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
              displayName: "Tenant Admins",
              entityType: "Group",
            },
          ],
        },
      ],
    },
  });
  console.log(result);
}
