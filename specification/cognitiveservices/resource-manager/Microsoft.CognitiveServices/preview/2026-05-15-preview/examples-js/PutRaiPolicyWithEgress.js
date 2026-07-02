const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update the state of specified Content Filters associated with the Azure OpenAI account.
 *
 * @summary update the state of specified Content Filters associated with the Azure OpenAI account.
 * x-ms-original-file: 2026-05-15-preview/PutRaiPolicyWithEgress.json
 */
async function putRaiPolicyWithEgress() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.raiPolicies.createOrUpdate(
    "resourceGroupName",
    "accountName",
    "egress-baseline",
    {
      properties: {
        basePolicyName: "Microsoft.Default",
        egressPolicy: {
          mode: "Enforced",
          defaultAction: "Deny",
          description: "Corporate baseline egress policy for sandboxed agents",
          rules: [
            {
              name: "allow-openai",
              description: "Allow traffic to OpenAI API",
              ruleType: "Fqdn",
              match: { host: "*.openai.com" },
              action: { actionType: "Allow" },
            },
            {
              name: "inject-auth-for-internal",
              description: "Inject managed identity token for internal services",
              ruleType: "Fqdn",
              match: { host: "*.internal.contoso.com" },
              action: {
                actionType: "Transform",
                headers: [
                  {
                    operation: "Set",
                    name: "Authorization",
                    valueRef: {
                      managedIdentityRef: {
                        resource: "https://internal.contoso.com/.default",
                        format: "Bearer {value}",
                      },
                    },
                  },
                ],
              },
            },
            {
              name: "rewrite-legacy-api",
              description: "Rewrite legacy API hostname to new internal endpoint",
              ruleType: "Fqdn",
              match: { host: "legacy-api.contoso.com", path: "/v1/*" },
              action: {
                actionType: "Rewrite",
                rewrite: { scheme: "https", host: "api-v2.internal.contoso.com", path: "/v2/" },
                headers: [
                  { operation: "Set", name: "X-Forwarded-Host", value: "legacy-api.contoso.com" },
                ],
              },
            },
          ],
        },
      },
    },
  );
  console.log(result);
}
