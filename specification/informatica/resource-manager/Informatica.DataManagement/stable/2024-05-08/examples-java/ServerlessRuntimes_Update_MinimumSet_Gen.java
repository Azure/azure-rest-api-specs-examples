
import com.azure.resourcemanager.informaticadatamanagement.models.InformaticaServerlessRuntimeResource;

/**
 * Samples for ServerlessRuntimes Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/
     * ServerlessRuntimes_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: ServerlessRuntimes_Update_Min.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void serverlessRuntimesUpdateMin(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        InformaticaServerlessRuntimeResource resource = manager.serverlessRuntimes()
            .getWithResponse("rgopenapi", "_f--", "8Zr__", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
