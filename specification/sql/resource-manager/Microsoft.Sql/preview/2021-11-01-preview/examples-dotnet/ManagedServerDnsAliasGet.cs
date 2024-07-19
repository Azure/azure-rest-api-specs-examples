using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/ManagedServerDnsAliasGet.json
// this example is just showing the usage of "ManagedServerDnsAliases_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedServerDnsAliasResource created on azure
// for more information of creating ManagedServerDnsAliasResource, please refer to the document of ManagedServerDnsAliasResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "Default";
string managedInstanceName = "dns-mi";
string dnsAliasName = "dns-alias-mi";
ResourceIdentifier managedServerDnsAliasResourceId = ManagedServerDnsAliasResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, dnsAliasName);
ManagedServerDnsAliasResource managedServerDnsAlias = client.GetManagedServerDnsAliasResource(managedServerDnsAliasResourceId);

// invoke the operation
ManagedServerDnsAliasResource result = await managedServerDnsAlias.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedServerDnsAliasData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
