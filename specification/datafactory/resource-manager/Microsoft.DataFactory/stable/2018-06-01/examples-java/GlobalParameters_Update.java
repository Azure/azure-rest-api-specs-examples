
import com.azure.resourcemanager.datafactory.models.GlobalParameterResource;

/**
 * Samples for GlobalParameters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * GlobalParameters_Update.json
     */
    /**
     * Sample code: GlobalParameters_Update.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void globalParametersUpdate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        GlobalParameterResource resource = manager.globalParameters()
            .getWithResponse("exampleResourceGroup", "exampleFactoryName", "default", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().apply();
    }
}
