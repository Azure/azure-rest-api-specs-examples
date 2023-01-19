using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Avs;
using Azure.ResourceManager.Avs.Models;

// Generated from example definition: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/ScriptExecutions_CreateOrUpdate.json
// this example is just showing the usage of "ScriptExecutions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AvsPrivateCloudResource created on azure
// for more information of creating AvsPrivateCloudResource, please refer to the document of AvsPrivateCloudResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
string privateCloudName = "cloud1";
ResourceIdentifier avsPrivateCloudResourceId = AvsPrivateCloudResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, privateCloudName);
AvsPrivateCloudResource avsPrivateCloud = client.GetAvsPrivateCloudResource(avsPrivateCloudResourceId);

// get the collection of this ScriptExecutionResource
ScriptExecutionCollection collection = avsPrivateCloud.GetScriptExecutions();

// invoke the operation
string scriptExecutionName = "addSsoServer";
ScriptExecutionData data = new ScriptExecutionData()
{
    ScriptCmdletId = new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/scriptPackages/AVS.PowerCommands@1.0.0/scriptCmdlets/New-SsoExternalIdentitySource"),
    Parameters =
    {
    new ScriptStringExecutionParameterDetails("DomainName")
    {
    Value = "placeholderDomain.local",
    },new ScriptStringExecutionParameterDetails("BaseUserDN")
    {
    Value = "DC=placeholder, DC=placeholder",
    }
    },
    HiddenParameters =
    {
    new ScriptSecureStringExecutionParameterDetails("Password")
    {
    SecureValue = "PlaceholderPassword",
    }
    },
    Timeout = "P0Y0M0DT0H60M60S",
    Retention = "P0Y0M60DT0H60M60S",
};
ArmOperation<ScriptExecutionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, scriptExecutionName, data);
ScriptExecutionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ScriptExecutionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
