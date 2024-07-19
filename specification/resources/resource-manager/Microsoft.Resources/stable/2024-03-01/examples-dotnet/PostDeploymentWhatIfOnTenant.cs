using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/PostDeploymentWhatIfOnTenant.json
// this example is just showing the usage of "Deployments_WhatIfAtTenantScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmDeploymentResource created on azure
// for more information of creating ArmDeploymentResource, please refer to the document of ArmDeploymentResource
string scope = $"/";
string deploymentName = "exampleDeploymentName";
ResourceIdentifier armDeploymentResourceId = ArmDeploymentResource.CreateResourceIdentifier(scope, deploymentName);
ArmDeploymentResource armDeployment = client.GetArmDeploymentResource(armDeploymentResourceId);

// invoke the operation
ArmDeploymentWhatIfContent content = new ArmDeploymentWhatIfContent(new ArmDeploymentWhatIfProperties(ArmDeploymentMode.Incremental)
{
    TemplateLink = new ArmDeploymentTemplateLink()
    {
        Uri = new Uri("https://example.com/exampleTemplate.json"),
    },
    Parameters = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
    }),
})
{
    Location = new AzureLocation("eastus"),
};
ArmOperation<WhatIfOperationResult> lro = await armDeployment.WhatIfAsync(WaitUntil.Completed, content);
WhatIfOperationResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
