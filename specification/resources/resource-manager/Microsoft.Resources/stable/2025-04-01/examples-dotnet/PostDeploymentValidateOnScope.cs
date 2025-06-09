using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2025-04-01/examples/PostDeploymentValidateOnScope.json
// this example is just showing the usage of "Deployments_ValidateAtScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmDeploymentResource created on azure
// for more information of creating ArmDeploymentResource, please refer to the document of ArmDeploymentResource
string scope = "subscriptions/00000000-0000-0000-0000-000000000001/resourceGroups/my-resource-group";
string deploymentName = "my-deployment";
ResourceIdentifier armDeploymentResourceId = ArmDeploymentResource.CreateResourceIdentifier(scope, deploymentName);
ArmDeploymentResource armDeployment = client.GetArmDeploymentResource(armDeploymentResourceId);

// invoke the operation
ArmDeploymentContent content = new ArmDeploymentContent(new ArmDeploymentProperties(ArmDeploymentMode.Incremental)
{
    TemplateLink = new ArmDeploymentTemplateLink
    {
        Uri = new Uri("https://example.com/exampleTemplate.json"),
        QueryString = "sv=2019-02-02&st=2019-04-29T22%3A18%3A26Z&se=2019-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=xxxxxxxx0xxxxxxxxxxxxx%2bxxxxxxxxxxxxxxxxxxxx%3d",
    },
    Parameters = BinaryData.FromObjectAsJson(new object()),
});
ArmOperation<ArmDeploymentValidateResult> lro = await armDeployment.ValidateAsync(WaitUntil.Completed, content);
ArmDeploymentValidateResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
