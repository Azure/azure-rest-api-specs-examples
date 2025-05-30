using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.LabServices.Models;
using Azure.ResourceManager.LabServices;

// Generated from example definition: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Users/patchUser.json
// this example is just showing the usage of "Users_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LabUserResource created on azure
// for more information of creating LabUserResource, please refer to the document of LabUserResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string labName = "testlab";
string userName = "testuser";
ResourceIdentifier labUserResourceId = LabUserResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, labName, userName);
LabUserResource labUser = client.GetLabUserResource(labUserResourceId);

// invoke the operation
LabUserPatch patch = new LabUserPatch
{
    AdditionalUsageQuota = XmlConvert.ToTimeSpan("PT10H"),
};
ArmOperation<LabUserResource> lro = await labUser.UpdateAsync(WaitUntil.Completed, patch);
LabUserResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LabUserData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
