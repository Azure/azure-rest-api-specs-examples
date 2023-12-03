using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/Compliances/GetCompliances_example.json
// this example is just showing the usage of "Compliances_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmResource created on azure
// for more information of creating ArmResource, please refer to the document of ArmResource

// get the collection of this SecurityComplianceResource
string scope = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23";
ResourceIdentifier scopeId = new ResourceIdentifier(string.Format("/{0}", scope));
SecurityComplianceCollection collection = client.GetSecurityCompliances(scopeId);

// invoke the operation and iterate over the result
await foreach (SecurityComplianceResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SecurityComplianceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
