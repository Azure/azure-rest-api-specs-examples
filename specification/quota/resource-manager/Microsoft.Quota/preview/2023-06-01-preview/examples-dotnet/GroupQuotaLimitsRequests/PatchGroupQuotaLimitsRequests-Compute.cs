using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Quota.Models;
using Azure.ResourceManager.Quota;

// Generated from example definition: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/GroupQuotaLimitsRequests/PatchGroupQuotaLimitsRequests-Compute.json
// this example is just showing the usage of "GroupQuotaLimitsRequest_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this GroupQuotaEntityResource created on azure
// for more information of creating GroupQuotaEntityResource, please refer to the document of GroupQuotaEntityResource
string managementGroupId = "E7EC67B3-7657-4966-BFFC-41EFD36BAA09";
string groupQuotaName = "groupquota1";
ResourceIdentifier groupQuotaEntityResourceId = GroupQuotaEntityResource.CreateResourceIdentifier(managementGroupId, groupQuotaName);
GroupQuotaEntityResource groupQuotaEntity = client.GetGroupQuotaEntityResource(groupQuotaEntityResourceId);

// invoke the operation
string resourceProviderName = "Microsoft.Compute";
string resourceName = "standardav2family";
ArmOperation<GroupQuotaRequestStatusResource> lro = await groupQuotaEntity.UpdateGroupQuotaLimitsRequestAsync(WaitUntil.Completed, resourceProviderName, resourceName);
GroupQuotaRequestStatusResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
GroupQuotaRequestStatusData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
