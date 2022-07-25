const { AzureDeploymentManager } = require("@azure/arm-deploymentmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Synchronously creates a new step or updates an existing step.
 *
 * @summary Synchronously creates a new step or updates an existing step.
 * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/step_health_check_createorupdate.json
 */
async function createHealthCheckStep() {
  const subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
  const resourceGroupName = "myResourceGroup";
  const stepName = "healthCheckStep";
  const stepInfo = {
    location: "centralus",
    properties: {
      attributes: {
        type: "REST",
        healthChecks: [
          {
            name: "appHealth",
            response: {
              regex: {
                matchQuantifier: "All",
                matches: [
                  "(?i)Contoso-App",
                  '(?i)"health_status":((.|\n)*)"(green|yellow)"',
                  '(?mi)^("application_host": 94781052)$',
                ],
              },
              successStatusCodes: ["OK"],
            },
            request: {
              method: "GET",
              authentication: {
                name: "Code",
                type: "ApiKey",
                in: "Query",
                value: "",
              },
              uri: "https://resthealth.healthservice.com/api/applications/contosoApp/healthStatus",
            },
          },
          {
            name: "serviceHealth",
            response: {
              regex: {
                matchQuantifier: "All",
                matches: ["(?i)Contoso-Service-EndToEnd", '(?i)"health_status":((.|\n)*)"(green)"'],
              },
              successStatusCodes: ["OK"],
            },
            request: {
              method: "GET",
              authentication: {
                name: "code",
                type: "ApiKey",
                in: "Header",
                value: "",
              },
              uri: "https://resthealth.healthservice.com/api/services/contosoService/healthStatus",
            },
          },
        ],
        healthyStateDuration: "PT2H",
        maxElasticDuration: "PT30M",
        waitDuration: "PT15M",
      },
      stepType: "HealthCheck",
    },
    tags: {},
  };
  const options = { stepInfo };
  const credential = new DefaultAzureCredential();
  const client = new AzureDeploymentManager(credential, subscriptionId);
  const result = await client.steps.createOrUpdate(resourceGroupName, stepName, options);
  console.log(result);
}

createHealthCheckStep().catch(console.error);
