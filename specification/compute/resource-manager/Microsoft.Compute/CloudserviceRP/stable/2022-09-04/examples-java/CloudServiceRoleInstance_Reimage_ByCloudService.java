
import com.azure.resourcemanager.compute.models.RoleInstances;
import java.util.Arrays;

/**
 * Samples for CloudServices Reimage.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/
     * CloudServiceRoleInstance_Reimage_ByCloudService.json
     */
    /**
     * Sample code: Reimage Cloud Service Role Instances in a Cloud Service.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        reimageCloudServiceRoleInstancesInACloudService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getCloudServices().reimage("ConstosoRG", "{cs-name}",
            new RoleInstances().withRoleInstances(Arrays.asList("ContosoFrontend_IN_0", "ContosoBackend_IN_1")),
            com.azure.core.util.Context.NONE);
    }
}
