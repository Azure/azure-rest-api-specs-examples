using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Workloads;
using Azure.ResourceManager.Workloads.Models;

// Generated from example definition: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/workloadmonitor/Db2ProviderInstances_Create_Root_Certificate.json
// this example is just showing the usage of "ProviderInstances_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SapMonitorResource created on azure
// for more information of creating SapMonitorResource, please refer to the document of SapMonitorResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "mySapMonitor";
ResourceIdentifier sapMonitorResourceId = SapMonitorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName);
SapMonitorResource sapMonitor = client.GetSapMonitorResource(sapMonitorResourceId);

// get the collection of this SapProviderInstanceResource
SapProviderInstanceCollection collection = sapMonitor.GetSapProviderInstances();

// invoke the operation
string providerInstanceName = "myProviderInstance";
SapProviderInstanceData data = new SapProviderInstanceData()
{
    ProviderSettings = new DB2ProviderInstanceProperties()
    {
        Hostname = "hostname",
        DBName = "dbName",
        DBPort = "dbPort",
        DBUsername = "username",
        DBPassword = "password",
        DBPasswordUri = new Uri(""),
        SapSid = "SID",
        SslPreference = SslPreference.RootCertificate,
    },
};
ArmOperation<SapProviderInstanceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, providerInstanceName, data);
SapProviderInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SapProviderInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
