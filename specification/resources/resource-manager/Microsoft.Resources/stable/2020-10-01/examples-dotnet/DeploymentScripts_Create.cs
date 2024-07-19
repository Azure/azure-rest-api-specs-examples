using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Resources/stable/2020-10-01/examples/DeploymentScripts_Create.json
// this example is just showing the usage of "DeploymentScripts_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "script-rg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ArmDeploymentScriptResource
ArmDeploymentScriptCollection collection = resourceGroupResource.GetArmDeploymentScripts();

// invoke the operation
string scriptName = "MyDeploymentScript";
ArmDeploymentScriptData data = new AzurePowerShellScript(new AzureLocation("westus"), XmlConvert.ToTimeSpan("PT7D"), "1.7.0")
{
    CleanupPreference = ScriptCleanupOptions.Always,
    SupportingScriptUris =
    {
    new Uri("https://uri1.to.supporting.script"),new Uri("https://uri2.to.supporting.script")
    },
    ScriptContent = "Param([string]$Location,[string]$Name) $deploymentScriptOutputs['test'] = 'value' Get-AzResourceGroup -Location $Location -Name $Name",
    Arguments = "-Location 'westus' -Name \"*rg2\"",
    Timeout = XmlConvert.ToTimeSpan("PT1H"),
    Identity = new ArmDeploymentScriptManagedIdentity()
    {
        IdentityType = ArmDeploymentScriptManagedIdentityType.UserAssigned,
        UserAssignedIdentities =
        {
        ["/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scriptRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/uai"] = new UserAssignedIdentity(),
        },
    },
};
ArmOperation<ArmDeploymentScriptResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, scriptName, data);
ArmDeploymentScriptResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ArmDeploymentScriptData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
