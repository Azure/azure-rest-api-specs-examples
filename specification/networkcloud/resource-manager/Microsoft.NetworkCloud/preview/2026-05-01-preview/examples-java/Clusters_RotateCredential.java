
import com.azure.resourcemanager.networkcloud.models.ClusterRotateCredentialParameters;
import java.util.Arrays;

/**
 * Samples for Clusters RotateCredential.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Clusters_RotateCredential.json
     */
    /**
     * Sample code: Rotate one or more managed credentials.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void
        rotateOneOrMoreManagedCredentials(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.clusters().rotateCredential("resourceGroupName", "clusterName",
            new ClusterRotateCredentialParameters().withCredentials(Arrays.asList("BMC Credential")),
            com.azure.core.util.Context.NONE);
    }
}
