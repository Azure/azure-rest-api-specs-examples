
import com.azure.resourcemanager.containerservice.fluent.models.TrustedAccessRoleBindingInner;
import java.util.Arrays;

/**
 * Samples for TrustedAccessRoleBindings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/TrustedAccessRoleBindings_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a trusted access role binding.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void createOrUpdateATrustedAccessRoleBinding(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getTrustedAccessRoleBindings().createOrUpdate("rg1", "clustername1", "binding1",
            new TrustedAccessRoleBindingInner().withSourceResourceId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/b/providers/Microsoft.MachineLearningServices/workspaces/c")
                .withRoles(Arrays.asList("Microsoft.MachineLearningServices/workspaces/reader",
                    "Microsoft.MachineLearningServices/workspaces/writer")),
            com.azure.core.util.Context.NONE);
    }
}
