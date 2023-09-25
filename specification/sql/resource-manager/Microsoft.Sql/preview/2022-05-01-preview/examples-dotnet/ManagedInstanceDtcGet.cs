using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;
using Azure.ResourceManager.Sql.Models;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-05-01-preview/examples/ManagedInstanceDtcGet.json
// this example is just showing the usage of "ManagedInstanceDtcs_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstanceDtcResource created on azure
// for more information of creating ManagedInstanceDtcResource, please refer to the document of ManagedInstanceDtcResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string managedInstanceName = "testinstance";
DtcName dtcName = DtcName.Current;
ResourceIdentifier managedInstanceDtcResourceId = ManagedInstanceDtcResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, dtcName);
ManagedInstanceDtcResource managedInstanceDtc = client.GetManagedInstanceDtcResource(managedInstanceDtcResourceId);

// invoke the operation
ManagedInstanceDtcResource result = await managedInstanceDtc.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedInstanceDtcData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
