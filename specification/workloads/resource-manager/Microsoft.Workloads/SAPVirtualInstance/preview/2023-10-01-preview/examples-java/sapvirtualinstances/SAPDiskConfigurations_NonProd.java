
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDatabaseType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDeploymentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDiskConfigurationsRequest;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapEnvironmentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapProductType;

/**
 * Samples for ResourceProvider SapDiskConfigurations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/
     * examples/sapvirtualinstances/SAPDiskConfigurations_NonProd.json
     */
    /**
     * Sample code: SAPDiskConfigurations_NonProd.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPDiskConfigurationsNonProd(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.resourceProviders().sapDiskConfigurationsWithResponse("centralus",
            new SapDiskConfigurationsRequest().withAppLocation("eastus").withEnvironment(SapEnvironmentType.NON_PROD)
                .withSapProduct(SapProductType.S4HANA).withDatabaseType(SapDatabaseType.HANA)
                .withDeploymentType(SapDeploymentType.THREE_TIER).withDbVmSku("Standard_M32ts"),
            com.azure.core.util.Context.NONE);
    }
}
