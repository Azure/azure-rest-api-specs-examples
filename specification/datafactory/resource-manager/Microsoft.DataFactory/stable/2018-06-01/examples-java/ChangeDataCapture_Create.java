import com.azure.resourcemanager.datafactory.models.MapperPolicy;
import com.azure.resourcemanager.datafactory.models.MapperSourceConnectionsInfo;
import com.azure.resourcemanager.datafactory.models.MapperTargetConnectionsInfo;
import java.util.List;

/** Samples for ChangeDataCapture CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ChangeDataCapture_Create.json
     */
    /**
     * Sample code: ChangeDataCapture_Create.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void changeDataCaptureCreate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .changeDataCaptures()
            .define("exampleChangeDataCapture")
            .withExistingFactory("exampleResourceGroup", "exampleFactoryName")
            .withSourceConnectionsInfo((List<MapperSourceConnectionsInfo>) null)
            .withTargetConnectionsInfo((List<MapperTargetConnectionsInfo>) null)
            .withPolicy((MapperPolicy) null)
            .withDescription(
                "Sample demo change data capture to transfer data from delimited (csv) to Azure SQL Database with"
                    + " automapped and non-automapped mappings.")
            .withAllowVNetOverride(false)
            .create();
    }
}
