using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HDInsight.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/HDI_Locations_ListBillingSpecs.json
// this example is just showing the usage of "Locations_ListBillingSpecs" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "subid";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("East US 2");
HDInsightBillingSpecsListResult result = await subscriptionResource.GetHDInsightBillingSpecsAsync(location);

Console.WriteLine($"Succeeded: {result}");
