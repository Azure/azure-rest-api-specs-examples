
import com.azure.resourcemanager.azurestackhci.models.SecretsLocationDetails;
import com.azure.resourcemanager.azurestackhci.models.SecretsLocationsChangeRequest;
import com.azure.resourcemanager.azurestackhci.models.SecretsType;
import java.util.Arrays;

/**
 * Samples for Clusters UpdateSecretsLocations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/Clusters_UpdateSecretsLocations.json
     */
    /**
     * Sample code: Update secrets locations for a Cluster.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        updateSecretsLocationsForACluster(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusters().updateSecretsLocations("test-rg", "myCluster",
            new SecretsLocationsChangeRequest().withProperties(Arrays.asList(new SecretsLocationDetails()
                .withSecretsType(SecretsType.BACKUP_SECRETS).withSecretsLocation("fakeTokenPlaceholder"))),
            com.azure.core.util.Context.NONE);
    }
}
