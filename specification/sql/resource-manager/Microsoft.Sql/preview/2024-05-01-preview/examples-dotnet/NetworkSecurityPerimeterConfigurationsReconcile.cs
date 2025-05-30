using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2024-05-01-preview/examples/NetworkSecurityPerimeterConfigurationsReconcile.json
// this example is just showing the usage of "NetworkSecurityPerimeterConfigurations_Reconcile" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SqlNetworkSecurityPerimeterConfigurationResource created on azure
// for more information of creating SqlNetworkSecurityPerimeterConfigurationResource, please refer to the document of SqlNetworkSecurityPerimeterConfigurationResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "sqlcrudtest-7398";
string serverName = "sqlcrudtest-7398";
string nspConfigName = "00000001-2222-3333-4444-111144444444.assoc1";
ResourceIdentifier sqlNetworkSecurityPerimeterConfigurationResourceId = SqlNetworkSecurityPerimeterConfigurationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serverName, nspConfigName);
SqlNetworkSecurityPerimeterConfigurationResource sqlNetworkSecurityPerimeterConfiguration = client.GetSqlNetworkSecurityPerimeterConfigurationResource(sqlNetworkSecurityPerimeterConfigurationResourceId);

// invoke the operation
ArmOperation<SqlNetworkSecurityPerimeterConfigurationResource> lro = await sqlNetworkSecurityPerimeterConfiguration.ReconcileAsync(WaitUntil.Completed);
SqlNetworkSecurityPerimeterConfigurationResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SqlNetworkSecurityPerimeterConfigurationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
