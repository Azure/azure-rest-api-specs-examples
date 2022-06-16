import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.loadtestservice.models.LoadTestResource;
import com.azure.resourcemanager.loadtestservice.models.LoadTestResourcePatchRequestBodyProperties;
import java.io.IOException;

/** Samples for LoadTests Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/preview/2021-12-01-preview/examples/LoadTests_Update.json
     */
    /**
     * Sample code: LoadTests_Update.
     *
     * @param manager Entry point to LoadTestManager.
     */
    public static void loadTestsUpdate(com.azure.resourcemanager.loadtestservice.LoadTestManager manager)
        throws IOException {
        LoadTestResource resource =
            manager.loadTests().getByResourceGroupWithResponse("dummyrg", "myLoadTest", Context.NONE).getValue();
        resource
            .update()
            .withTags(
                SerializerFactory
                    .createDefaultManagementSerializerAdapter()
                    .deserialize("{\"Division\":\"LT\",\"Team\":\"Dev Exp\"}", Object.class, SerializerEncoding.JSON))
            .withProperties(
                new LoadTestResourcePatchRequestBodyProperties().withDescription("This is new load test resource"))
            .apply();
    }
}
