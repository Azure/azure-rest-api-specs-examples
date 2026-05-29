const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a governance rule over a given scope
 *
 * @summary creates or updates a governance rule over a given scope
 * x-ms-original-file: 2022-01-01-preview/GovernanceRules/PutSecurityConnectorGovernanceRule_example.json
 */
async function createOrUpdateGovernanceRuleOverSecurityConnectorScope() {
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.governanceRules.createOrUpdate(
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/gcpResourceGroup/providers/Microsoft.Security/securityConnectors/gcpconnector",
    "ad9a8e26-29d9-4829-bb30-e597a58cdbb8",
    {
      description: "A rule on critical GCP recommendations",
      conditionSets: [
        {
          conditions: [
            {
              operator: "In",
              property: "$.AssessmentKey",
              value:
                '["b1cd27e0-4ecc-4246-939f-49c426d9d72f", "fe83f80b-073d-4ccf-93d9-6797eb870201"]',
            },
          ],
        },
      ],
      displayName: "GCP Admin's rule",
      governanceEmailNotification: {
        disableManagerEmailNotification: true,
        disableOwnerEmailNotification: false,
      },
      isDisabled: false,
      isGracePeriod: true,
      ownerSource: { type: "Manually", value: "user@contoso.com" },
      remediationTimeframe: "7.00:00:00",
      rulePriority: 200,
      ruleType: "Integrated",
      sourceResourceType: "Assessments",
    },
  );
  console.log(result);
}
