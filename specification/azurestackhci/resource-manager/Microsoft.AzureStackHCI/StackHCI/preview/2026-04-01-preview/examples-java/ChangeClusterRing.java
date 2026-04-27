
import com.azure.resourcemanager.azurestackhci.models.ChangeRingRequest;
import com.azure.resourcemanager.azurestackhci.models.ChangeRingRequestProperties;

/**
 * Samples for Clusters ChangeRing.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/ChangeClusterRing.json
     */
    /**
     * Sample code: Change cluster ring.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void changeClusterRing(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusters().changeRing("test-rg", "myCluster",
            new ChangeRingRequest().withProperties(new ChangeRingRequestProperties().withTargetRing("Insider")),
            com.azure.core.util.Context.NONE);
    }
}
