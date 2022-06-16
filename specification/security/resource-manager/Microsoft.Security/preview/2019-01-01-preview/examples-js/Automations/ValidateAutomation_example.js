const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Validates the security automation model before create or update. Any validation errors are returned to the client.
 *
 * @summary Validates the security automation model before create or update. Any validation errors are returned to the client.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/ValidateAutomation_example.json
 */
async function validateTheSecurityAutomationModelBeforeCreateOrUpdate() {
  const subscriptionId = "a5caac9c-5c04-49af-b3d0-e204f40345d5";
  const resourceGroupName = "exampleResourceGroup";
  const automationName = "exampleAutomation";
  const automation = {
    description:
      "An example of a security automation that triggers one LogicApp resource (myTest1) on any security assessment of type customAssessment",
    actions: [
      {
        actionType: "LogicApp",
        logicAppResourceId:
          "/subscriptions/e54a4a18-5b94-4f90-9471-bd3decad8a2e/resourceGroups/sample/providers/Microsoft.Logic/workflows/MyTest1",
        uri: "https://exampleTriggerUri1.com",
      },
    ],
    isEnabled: true,
    location: "Central US",
    scopes: [
      {
        description:
          "A description that helps to identify this scope - for example: security assessments that relate to the resource group myResourceGroup within the subscription a5caac9c-5c04-49af-b3d0-e204f40345d5",
        scopePath:
          "/subscriptions/a5caac9c-5c04-49af-b3d0-e204f40345d5/resourceGroups/myResourceGroup",
      },
    ],
    sources: [
      {
        eventSource: "Assessments",
        ruleSets: [
          {
            rules: [
              {
                expectedValue: "customAssessment",
                operator: "Equals",
                propertyJPath: "$.Entity.AssessmentType",
                propertyType: "String",
              },
            ],
          },
        ],
      },
    ],
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.automations.validate(resourceGroupName, automationName, automation);
  console.log(result);
}

validateTheSecurityAutomationModelBeforeCreateOrUpdate().catch(console.error);
