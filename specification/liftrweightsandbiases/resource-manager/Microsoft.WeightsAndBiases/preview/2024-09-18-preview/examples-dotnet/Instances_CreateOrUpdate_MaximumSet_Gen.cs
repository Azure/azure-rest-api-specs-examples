using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.WeightsAndBiases.Models;
using Azure.ResourceManager.WeightsAndBiases;

// Generated from example definition: 2024-09-18-preview/Instances_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "InstanceResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "0BCB047F-CB71-4DFE-B0BD-88519F411C2F";
string resourceGroupName = "rgopenapi";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this WeightsAndBiasesInstanceResource
WeightsAndBiasesInstanceCollection collection = resourceGroupResource.GetWeightsAndBiasesInstances();

// invoke the operation
string instancename = "myinstance";
WeightsAndBiasesInstanceData data = new WeightsAndBiasesInstanceData(new AzureLocation("pudewmshbcvbt"))
{
    Properties = new WeightsAndBiasesInstanceProperties(new WeightsAndBiasesMarketplaceDetails(new WeightsAndBiasesOfferDetails("kf", "rfgoevxeke", "ufopn")
    {
        PlanName = "adysakgqlryufffz",
        TermUnit = "dgrkojow",
        TermId = "kklscqq",
    })
    {
        SubscriptionId = "00000000-0000-0000-0000-000000000000",
    }, new WeightsAndBiasesUserDetails
    {
        FirstName = "kiiehcojcldrlndoid",
        LastName = "zhkvsfqvthwkfkvgxcruyud",
        EmailAddress = "user@outlook.com",
        Upn = "rmjpgqchpbw",
        PhoneNumber = "cogmqmuwfcpstkwbzgkgo",
    })
    {
        PartnerProperties = new WeightsAndBiasesPartnerProperties(WeightsAndBiasesRegion.EastUS, "xkecokwnwtkwnkxfgucmzybzzb"),
        SingleSignOnProperties = new WeightsAndBiasesSingleSignOnPropertiesV2(WeightsAndBiasesSingleSignOnType.Saml)
        {
            State = WeightsAndBiasesSingleSignOnState.Initial,
            EnterpriseAppId = "hkxtmpv",
            Uri = "iqlemoksqdygqyxpp",
            AadDomains = { "mxnw" },
        },
    },
    Identity = new ManagedServiceIdentity("None")
    {
        UserAssignedIdentities = { },
    },
    Tags = { },
};
ArmOperation<WeightsAndBiasesInstanceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, instancename, data);
WeightsAndBiasesInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WeightsAndBiasesInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
