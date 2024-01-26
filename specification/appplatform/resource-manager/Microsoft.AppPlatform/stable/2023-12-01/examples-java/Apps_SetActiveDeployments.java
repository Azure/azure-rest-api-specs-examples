
import com.azure.resourcemanager.appplatform.models.ActiveDeploymentCollection;
import java.util.Arrays;

/**
 * Samples for Apps SetActiveDeployments.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * Apps_SetActiveDeployments.json
     */
    /**
     * Sample code: Apps_SetActiveDeployments.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void appsSetActiveDeployments(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getApps().setActiveDeployments("myResourceGroup", "myservice",
            "myapp", new ActiveDeploymentCollection().withActiveDeploymentNames(Arrays.asList("default")),
            com.azure.core.util.Context.NONE);
    }
}
