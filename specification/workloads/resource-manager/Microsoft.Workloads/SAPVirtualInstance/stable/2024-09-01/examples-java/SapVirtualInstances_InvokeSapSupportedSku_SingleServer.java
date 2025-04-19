
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDatabaseType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapDeploymentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapEnvironmentType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapProductType;
import com.azure.resourcemanager.workloadssapvirtualinstance.models.SapSupportedSkusRequest;

/**
 * Samples for SapVirtualInstances GetSapSupportedSku.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/SapVirtualInstances_InvokeSapSupportedSku_SingleServer.json
     */
    /**
     * Sample code: SAP supported SKUs for single server.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void sAPSupportedSKUsForSingleServer(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapVirtualInstances().getSapSupportedSkuWithResponse("centralus",
            new SapSupportedSkusRequest().withAppLocation("eastus").withEnvironment(SapEnvironmentType.NON_PROD)
                .withSapProduct(SapProductType.S4HANA).withDeploymentType(SapDeploymentType.SINGLE_SERVER)
                .withDatabaseType(SapDatabaseType.HANA),
            com.azure.core.util.Context.NONE);
    }
}
