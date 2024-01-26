
import com.azure.resourcemanager.appservice.models.KubeEnvironmentPatchResource;

/**
 * Samples for KubeEnvironments Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/KubeEnvironments_Update.json
     */
    /**
     * Sample code: Update kube environments.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateKubeEnvironments(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getKubeEnvironments().updateWithResponse("examplerg", "testkubeenv",
            new KubeEnvironmentPatchResource().withStaticIp("1.2.3.4"), com.azure.core.util.Context.NONE);
    }
}
