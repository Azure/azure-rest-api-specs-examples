using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PostDeploymentWhatIfOnSubscription.json
// this example is just showing the usage of "Deployments_WhatIfAtSubscriptionScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmDeploymentResource created on azure
// for more information of creating ArmDeploymentResource, please refer to the document of ArmDeploymentResource
string subscriptionId = "00000000-0000-0000-0000-000000000001";
string scope = $"/subscriptions/{subscriptionId}";
string deploymentName = "my-deployment";
ResourceIdentifier armDeploymentResourceId = ArmDeploymentResource.CreateResourceIdentifier(scope, deploymentName);
ArmDeploymentResource armDeployment = client.GetArmDeploymentResource(armDeploymentResourceId);

// invoke the operation
ArmDeploymentWhatIfContent content = new ArmDeploymentWhatIfContent(new ArmDeploymentWhatIfProperties(ArmDeploymentMode.Incremental)
{
    TemplateLink = null,
    Parameters = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
    }),
})
{
    Location = new AzureLocation("westus"),
};
ArmOperation<WhatIfOperationResult> lro = await armDeployment.WhatIfAsync(WaitUntil.Completed, content);
WhatIfOperationResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
