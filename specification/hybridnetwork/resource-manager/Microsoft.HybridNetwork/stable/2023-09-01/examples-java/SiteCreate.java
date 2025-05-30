
import com.azure.resourcemanager.hybridnetwork.models.AzureArcK8SClusterNfviDetails;
import com.azure.resourcemanager.hybridnetwork.models.AzureCoreNfviDetails;
import com.azure.resourcemanager.hybridnetwork.models.AzureOperatorNexusClusterNfviDetails;
import com.azure.resourcemanager.hybridnetwork.models.ReferencedResource;
import com.azure.resourcemanager.hybridnetwork.models.SitePropertiesFormat;
import java.util.Arrays;

/**
 * Samples for Sites CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/SiteCreate.json
     */
    /**
     * Sample code: Create network site.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void createNetworkSite(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.sites().define("testSite").withRegion("westUs2").withExistingResourceGroup("rg1")
            .withProperties(new SitePropertiesFormat().withNfvis(Arrays.asList(
                new AzureCoreNfviDetails().withName("nfvi1").withLocation("westUs2"),
                new AzureArcK8SClusterNfviDetails().withName("nfvi2")
                    .withCustomLocationReference(new ReferencedResource().withId(
                        "/subscriptions/subid/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation1")),
                new AzureOperatorNexusClusterNfviDetails().withName("nfvi3")
                    .withCustomLocationReference(new ReferencedResource().withId(
                        "/subscriptions/subid/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation2")))))
            .create();
    }
}
