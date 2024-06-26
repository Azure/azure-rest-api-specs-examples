const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Organization role bindings
 *
 * @summary Organization role bindings
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Access_RoleBindingNameList.json
 */
async function accessRoleBindingNameList() {
  const subscriptionId =
    process.env["CONFLUENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONFLUENT_RESOURCE_GROUP"] || "myResourceGroup";
  const organizationName = "myOrganization";
  const body = {
    searchFilters: {
      crnPattern: "crn://confluent.cloud/organization=1aa7de07-298e-479c-8f2f-16ac91fd8e76",
      namespace:
        "public,dataplane,networking,identity,datagovernance,connect,streamcatalog,pipelines,ksql",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.access.listRoleBindingNameList(
    resourceGroupName,
    organizationName,
    body,
  );
  console.log(result);
}
