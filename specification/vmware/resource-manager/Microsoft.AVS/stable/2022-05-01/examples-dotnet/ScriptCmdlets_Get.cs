using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Avs;

// Generated from example definition: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/ScriptCmdlets_Get.json
// this example is just showing the usage of "ScriptCmdlets_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ScriptPackageResource created on azure
// for more information of creating ScriptPackageResource, please refer to the document of ScriptPackageResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "{privateCloudName}";
string scriptPackageName = "{scriptPackageName}";
ResourceIdentifier scriptPackageResourceId = ScriptPackageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName, scriptPackageName);
ScriptPackageResource scriptPackage = client.GetScriptPackageResource(scriptPackageResourceId);

// get the collection of this ScriptCmdletResource
ScriptCmdletCollection collection = scriptPackage.GetScriptCmdlets();

// invoke the operation
string scriptCmdletName = "New-ExternalSsoDomain";
bool result = await collection.ExistsAsync(scriptCmdletName);

Console.WriteLine($"Succeeded: {result}");
