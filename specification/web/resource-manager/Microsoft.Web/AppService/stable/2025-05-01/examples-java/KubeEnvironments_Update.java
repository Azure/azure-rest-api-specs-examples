
import com.azure.resourcemanager.appservice.models.KubeEnvironmentPatchResource;

/**
 * Samples for KubeEnvironments Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/KubeEnvironments_Update.json
     */
    /**
     * Sample code: Update kube environments.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void updateKubeEnvironments(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getKubeEnvironments().updateWithResponse("examplerg", "testkubeenv",
            new KubeEnvironmentPatchResource().withStaticIp("1.2.3.4"), com.azure.core.util.Context.NONE);
    }
}
