using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2023-08-01/examples/DeploymentScripts_Get.json
// this example is just showing the usage of "DeploymentScripts_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmDeploymentScriptResource created on azure
// for more information of creating ArmDeploymentScriptResource, please refer to the document of ArmDeploymentScriptResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "script-rg";
string scriptName = "MyDeploymentScript";
ResourceIdentifier armDeploymentScriptResourceId = ArmDeploymentScriptResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, scriptName);
ArmDeploymentScriptResource armDeploymentScript = client.GetArmDeploymentScriptResource(armDeploymentScriptResourceId);

// invoke the operation
ArmDeploymentScriptResource result = await armDeploymentScript.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ArmDeploymentScriptData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
