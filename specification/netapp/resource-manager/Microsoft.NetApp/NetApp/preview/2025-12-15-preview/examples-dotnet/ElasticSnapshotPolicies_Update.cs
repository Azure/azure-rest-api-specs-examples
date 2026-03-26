using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-12-15-preview/examples/ElasticSnapshotPolicies_Update.json
// this example is just showing the usage of "ElasticSnapshotPolicies_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppElasticSnapshotPolicyResource created on azure
// for more information of creating NetAppElasticSnapshotPolicyResource, please refer to the document of NetAppElasticSnapshotPolicyResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string snapshotPolicyName = "snapshotPolicyName";
ResourceIdentifier netAppElasticSnapshotPolicyResourceId = NetAppElasticSnapshotPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, snapshotPolicyName);
NetAppElasticSnapshotPolicyResource netAppElasticSnapshotPolicy = client.GetNetAppElasticSnapshotPolicyResource(netAppElasticSnapshotPolicyResourceId);

// invoke the operation
NetAppElasticSnapshotPolicyPatch patch = new NetAppElasticSnapshotPolicyPatch
{
    Properties = new NetAppElasticSnapshotPolicyPatchProperties
    {
        HourlySchedule = new ElasticSnapshotPolicyHourlySchedule
        {
            SnapshotsToKeep = 2,
            Minute = 50,
        },
        DailySchedule = new ElasticSnapshotPolicyDailySchedule
        {
            SnapshotsToKeep = 4,
            Hour = 14,
            Minute = 30,
        },
        WeeklySchedule = new ElasticSnapshotPolicyWeeklySchedule
        {
            SnapshotsToKeep = 3,
            Days = { NetAppDayOfWeek.Wednesday },
            Hour = 14,
            Minute = 45,
        },
        MonthlySchedule = new ElasticSnapshotPolicyMonthlySchedule
        {
            SnapshotsToKeep = 5,
            DaysOfMonth = { 10, 11, 12 },
            Hour = 14,
            Minute = 15,
        },
        PolicyStatus = NetAppPolicyStatus.Enabled,
    },
};
ArmOperation<NetAppElasticSnapshotPolicyResource> lro = await netAppElasticSnapshotPolicy.UpdateAsync(WaitUntil.Completed, patch);
NetAppElasticSnapshotPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetAppElasticSnapshotPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
