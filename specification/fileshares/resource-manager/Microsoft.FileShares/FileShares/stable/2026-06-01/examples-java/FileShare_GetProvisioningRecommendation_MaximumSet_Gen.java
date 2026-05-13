
import com.azure.resourcemanager.fileshares.models.FileShareProvisioningRecommendationInput;
import com.azure.resourcemanager.fileshares.models.FileShareProvisioningRecommendationRequest;

/**
 * Samples for InformationalOperations GetProvisioningRecommendation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/FileShare_GetProvisioningRecommendation_MaximumSet_Gen.json
     */
    /**
     * Sample code: FileShare_GetProvisioningRecommendation_MaximumSet.
     * 
     * @param manager Entry point to FileSharesManager.
     */
    public static void fileShareGetProvisioningRecommendationMaximumSet(
        com.azure.resourcemanager.fileshares.FileSharesManager manager) {
        manager.informationalOperations().getProvisioningRecommendationWithResponse("westus",
            new FileShareProvisioningRecommendationRequest()
                .withProperties(new FileShareProvisioningRecommendationInput().withProvisionedStorageGiB(7)),
            com.azure.core.util.Context.NONE);
    }
}
