using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/stable/2023-08-01/examples/ManagedInstanceDtcUpdateMax.json
// this example is just showing the usage of "ManagedInstanceDtcs_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstanceResource created on azure
// for more information of creating ManagedInstanceResource, please refer to the document of ManagedInstanceResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "testrg";
string managedInstanceName = "testinstance";
ResourceIdentifier managedInstanceResourceId = ManagedInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName);
ManagedInstanceResource managedInstance = client.GetManagedInstanceResource(managedInstanceResourceId);

// get the collection of this ManagedInstanceDtcResource
ManagedInstanceDtcCollection collection = managedInstance.GetManagedInstanceDtcs();

// invoke the operation
DtcName dtcName = DtcName.Current;
ManagedInstanceDtcData data = new ManagedInstanceDtcData
{
    DtcEnabled = true,
    SecuritySettings = new ManagedInstanceDtcSecuritySettings
    {
        TransactionManagerCommunicationSettings = new ManagedInstanceDtcTransactionManagerCommunicationSettings
        {
            AllowInboundEnabled = false,
            AllowOutboundEnabled = true,
            Authentication = "NoAuth",
        },
        IsXATransactionsEnabled = false,
        SnaLu6Point2TransactionsEnabled = false,
        XATransactionsDefaultTimeoutInSeconds = 1000,
        XATransactionsMaximumTimeoutInSeconds = 3000,
    },
    ExternalDnsSuffixSearchList = { "dns.example1.com", "dns.example2.com" },
};
ArmOperation<ManagedInstanceDtcResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, dtcName, data);
ManagedInstanceDtcResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedInstanceDtcData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
