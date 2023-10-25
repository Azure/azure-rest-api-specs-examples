const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing delivery rule within a rule set.
 *
 * @summary Updates an existing delivery rule within a rule set.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Rules_Update.json
 */
async function rulesUpdate() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const ruleSetName = "ruleSet1";
  const ruleName = "rule1";
  const ruleUpdateProperties = {
    actions: [
      {
        name: "ModifyResponseHeader",
        parameters: {
          headerAction: "Overwrite",
          headerName: "X-CDN",
          typeName: "DeliveryRuleHeaderActionParameters",
          value: "MSFT",
        },
      },
    ],
    order: 1,
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.rules.beginUpdateAndWait(
    resourceGroupName,
    profileName,
    ruleSetName,
    ruleName,
    ruleUpdateProperties
  );
  console.log(result);
}
