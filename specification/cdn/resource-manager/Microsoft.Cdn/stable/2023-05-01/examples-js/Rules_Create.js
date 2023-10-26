const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new delivery rule within the specified rule set.
 *
 * @summary Creates a new delivery rule within the specified rule set.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Rules_Create.json
 */
async function rulesCreate() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const ruleSetName = "ruleSet1";
  const ruleName = "rule1";
  const rule = {
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
    conditions: [
      {
        name: "RequestMethod",
        parameters: {
          matchValues: ["GET"],
          negateCondition: false,
          operator: "Equal",
          typeName: "DeliveryRuleRequestMethodConditionParameters",
        },
      },
    ],
    order: 1,
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.rules.beginCreateAndWait(
    resourceGroupName,
    profileName,
    ruleSetName,
    ruleName,
    rule
  );
  console.log(result);
}
