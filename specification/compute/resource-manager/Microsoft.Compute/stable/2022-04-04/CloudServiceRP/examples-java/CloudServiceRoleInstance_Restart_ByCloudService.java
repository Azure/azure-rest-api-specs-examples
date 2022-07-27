import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.RoleInstances;
import java.util.Arrays;

/** Samples for CloudServices Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-04-04/CloudServiceRP/examples/CloudServiceRoleInstance_Restart_ByCloudService.json
     */
    /**
     * Sample code: Restart Cloud Service Role Instances in a Cloud Service.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void restartCloudServiceRoleInstancesInACloudService(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getCloudServices()
            .restart(
                "ConstosoRG",
                "{cs-name}",
                new RoleInstances().withRoleInstances(Arrays.asList("ContosoFrontend_IN_0", "ContosoBackend_IN_1")),
                Context.NONE);
    }
}
