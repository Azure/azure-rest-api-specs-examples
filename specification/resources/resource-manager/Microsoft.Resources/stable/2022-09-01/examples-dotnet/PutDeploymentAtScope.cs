using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2022-09-01/examples/PutDeploymentAtScope.json
// this example is just showing the usage of "Deployments_CreateOrUpdateAtScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmDeploymentResource created on azure
// for more information of creating ArmDeploymentResource, please refer to the document of ArmDeploymentResource
string scope = "providers/Microsoft.Management/managementGroups/my-management-group-id";
string deploymentName = "my-deployment";
ResourceIdentifier armDeploymentResourceId = ArmDeploymentResource.CreateResourceIdentifier(scope, deploymentName);
ArmDeploymentResource armDeployment = client.GetArmDeploymentResource(armDeploymentResourceId);

// invoke the operation
ArmDeploymentContent content = new ArmDeploymentContent(new ArmDeploymentProperties(ArmDeploymentMode.Incremental)
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
    Tags =
    {
    ["tagKey1"] = "tag-value-1",
    ["tagKey2"] = "tag-value-2",
    },
};
ArmOperation<ArmDeploymentResource> lro = await armDeployment.UpdateAsync(WaitUntil.Completed, content);
ArmDeploymentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ArmDeploymentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
