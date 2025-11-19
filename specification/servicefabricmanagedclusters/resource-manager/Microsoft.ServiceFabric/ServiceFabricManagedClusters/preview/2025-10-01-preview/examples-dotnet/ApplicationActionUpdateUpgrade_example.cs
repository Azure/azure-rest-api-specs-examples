using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceFabricManagedClusters.Models;
using Azure.ResourceManager.ServiceFabricManagedClusters;

// Generated from example definition: 2025-10-01-preview/ApplicationActionUpdateUpgrade_example.json
// this example is just showing the usage of "Applications_UpdateUpgrade" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceFabricManagedApplicationResource created on azure
// for more information of creating ServiceFabricManagedApplicationResource, please refer to the document of ServiceFabricManagedApplicationResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string clusterName = "myCluster";
string applicationName = "myApp";
ResourceIdentifier serviceFabricManagedApplicationResourceId = ServiceFabricManagedApplicationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, applicationName);
ServiceFabricManagedApplicationResource serviceFabricManagedApplication = client.GetServiceFabricManagedApplicationResource(serviceFabricManagedApplicationResourceId);

// invoke the operation
RuntimeUpdateApplicationUpgradeContent content = new RuntimeUpdateApplicationUpgradeContent("fabric:/Voting", RuntimeUpgradeKind.Rolling)
{
    ApplicationHealthPolicy = new RuntimeApplicationHealthPolicy(true, 10)
    {
        DefaultServiceTypeHealthPolicy = new RuntimeServiceTypeHealthPolicy(12, 10, 11),
        ServiceTypeHealthPolicyMap =
        {
        ["VotingWeb"] = new RuntimeServiceTypeHealthPolicy(15, 13, 14)
        },
    },
    UpdateDescription = new RuntimeRollingUpgradeUpdateMonitoringPolicy(RuntimeRollingUpgradeMode.Monitored)
    {
        ForceRestart = true,
        FailureAction = RuntimeFailureAction.Manual,
        HealthCheckWaitDurationInMilliseconds = "PT0H0M10S",
        HealthCheckStableDurationInMilliseconds = "PT1H0M0S",
        HealthCheckRetryTimeoutInMilliseconds = "PT0H15M0S",
        UpgradeTimeoutInMilliseconds = "PT2H0M0S",
        UpgradeDomainTimeoutInMilliseconds = "PT2H0M0S",
    },
};
await serviceFabricManagedApplication.UpdateUpgradeAsync(WaitUntil.Completed, content);

Console.WriteLine("Succeeded");
