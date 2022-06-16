import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.fluent.models.KubeEnvironmentInner;

/** Samples for KubeEnvironments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/KubeEnvironments_CreateOrUpdate.json
     */
    /**
     * Sample code: Create kube environments.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createKubeEnvironments(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getKubeEnvironments()
            .createOrUpdate(
                "examplerg",
                "testkubeenv",
                new KubeEnvironmentInner().withLocation("East US").withStaticIp("1.2.3.4"),
                Context.NONE);
    }
}
