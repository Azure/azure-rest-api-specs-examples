
import com.azure.resourcemanager.providerhub.models.ResourceManagementAction;
import com.azure.resourcemanager.providerhub.models.ResourceManagementEntity;
import java.util.Arrays;

/**
 * Samples for ResourceActions DeleteResources.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * ResourceActions_DeleteResources.json
     */
    /**
     * Sample code: ResourceActions_DeleteResources.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void
        resourceActionsDeleteResources(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.resourceActions().deleteResources("Microsoft.Contoso", "default",
            new ResourceManagementAction().withResources(Arrays.asList(new ResourceManagementEntity()
                .withResourceId(
                    "/subscriptions/ab7a8701-f7ef-471a-a2f4-d0ebbf494f77/providers/Microsoft.Contoso/employee/test")
                .withHomeTenantId("11111111-f7ef-471a-a2f4-d0ebbf494f77").withLocation("southeastasia"))),
            com.azure.core.util.Context.NONE);
    }
}
