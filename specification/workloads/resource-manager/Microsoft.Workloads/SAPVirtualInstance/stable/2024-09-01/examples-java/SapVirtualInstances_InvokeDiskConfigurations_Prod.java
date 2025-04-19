
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDatabaseType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDeploymentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDiskConfigurationsRequest;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapEnvironmentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapProductType;

/**
 * Samples for SapVirtualInstances GetDiskConfigurations.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_InvokeDiskConfigurations_Prod.json
     */
    /**
     * Sample code: SAP disk configurations for input environment Prod.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPDiskConfigurationsForInputEnvironmentProd(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().getDiskConfigurationsWithResponse("centralus",
            new SapDiskConfigurationsRequest().withAppLocation("eastus").withEnvironment(SapEnvironmentType.PROD)
                .withSapProduct(SapProductType.S4HANA).withDatabaseType(SapDatabaseType.HANA)
                .withDeploymentType(SapDeploymentType.THREE_TIER).withDbVmSku("Standard_M32ts"),
            com.azure.core.util.Context.NONE);
    }
}
