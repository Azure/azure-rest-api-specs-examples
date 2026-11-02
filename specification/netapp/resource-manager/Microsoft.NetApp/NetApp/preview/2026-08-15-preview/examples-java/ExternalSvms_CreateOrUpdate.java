
import com.azure.resourcemanager.netapp.models.ExternalSvmProperties;

/**
 * Samples for ExternalSvms CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ExternalSvms_CreateOrUpdate.json
     */
    /**
     * Sample code: ExternalSvms_CreateOrUpdate.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void externalSvmsCreateOrUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.externalSvms().define("contoso-svm").withRegion("eastus")
            .withExistingExternalCluster("contosoNetAppRg", "contoso-external-cluster")
            .withProperties(new ExternalSvmProperties().withOntapSvmName("contoso-svm")).create();
    }
}
