const { AzureDeploymentManager } = require("@azure/arm-deploymentmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This is an asynchronous operation and can be polled to completion using the location header returned by this operation.
 *
 * @summary This is an asynchronous operation and can be polled to completion using the location header returned by this operation.
 * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/rollout_createorupdate.json
 */
async function createOrUpdateRollout() {
  const subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
  const resourceGroupName = "myResourceGroup";
  const rolloutName = "myRollout";
  const rolloutRequest = {
    artifactSourceId:
      "/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/Microsoft.DeploymentManager/artifactSources/myArtifactSource",
    buildVersion: "1.0.0.1",
    identity: {
      type: "userAssigned",
      identityIds: [
        "/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userassignedidentities/myuseridentity",
      ],
    },
    location: "centralus",
    stepGroups: [
      {
        name: "FirstRegion",
        deploymentTargetId:
          "Microsoft.DeploymentManager/serviceTopologies/myTopology/services/myService/serviceUnits/myServiceUnit1'",
        postDeploymentSteps: [{ stepId: "Microsoft.DeploymentManager/steps/postDeployStep1" }],
        preDeploymentSteps: [
          { stepId: "Microsoft.DeploymentManager/steps/preDeployStep1" },
          { stepId: "Microsoft.DeploymentManager/steps/preDeployStep2" },
        ],
      },
      {
        name: "SecondRegion",
        dependsOnStepGroups: ["FirstRegion"],
        deploymentTargetId:
          "Microsoft.DeploymentManager/serviceTopologies/myTopology/services/myService/serviceUnits/myServiceUnit2'",
        postDeploymentSteps: [{ stepId: "Microsoft.DeploymentManager/steps/postDeployStep5" }],
        preDeploymentSteps: [
          { stepId: "Microsoft.DeploymentManager/steps/preDeployStep3" },
          { stepId: "Microsoft.DeploymentManager/steps/preDeployStep4" },
        ],
      },
    ],
    tags: {},
    targetServiceTopologyId:
      "/subscriptions/caac1590-e859-444f-a9e0-62091c0f5929/resourceGroups/myResourceGroup/Microsoft.DeploymentManager/serviceTopologies/myTopology",
  };
  const options = { rolloutRequest };
  const credential = new DefaultAzureCredential();
  const client = new AzureDeploymentManager(credential, subscriptionId);
  const result = await client.rollouts.beginCreateOrUpdateAndWait(
    resourceGroupName,
    rolloutName,
    options
  );
  console.log(result);
}

createOrUpdateRollout().catch(console.error);
