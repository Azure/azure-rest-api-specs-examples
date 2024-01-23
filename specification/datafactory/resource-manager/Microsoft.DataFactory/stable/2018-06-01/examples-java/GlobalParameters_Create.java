
import com.azure.resourcemanager.datafactory.models.GlobalParameterSpecification;
import java.util.Map;

/**
 * Samples for GlobalParameters CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * GlobalParameters_Create.json
     */
    /**
     * Sample code: GlobalParameters_Create.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void globalParametersCreate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.globalParameters().define("default").withExistingFactory("exampleResourceGroup", "exampleFactoryName")
            .withProperties((Map<String, GlobalParameterSpecification>) null).create();
    }
}
