import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.servicenetworking.models.TrafficController;
import java.io.IOException;

/** Samples for TrafficControllerInterface Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/TrafficControllerPatch.json
     */
    /**
     * Sample code: Patch Traffic Controller.
     *
     * @param manager Entry point to TrafficControllerManager.
     */
    public static void patchTrafficController(
        com.azure.resourcemanager.servicenetworking.TrafficControllerManager manager) throws IOException {
        TrafficController resource =
            manager.trafficControllerInterfaces().getByResourceGroupWithResponse("rg1", "TC1", Context.NONE).getValue();
        resource
            .update()
            .withProperties(
                SerializerFactory
                    .createDefaultManagementSerializerAdapter()
                    .deserialize(
                        "{\"configurationEndpoints\":[\"abc.eastus.trafficcontroller.azure.net\"]}",
                        Object.class,
                        SerializerEncoding.JSON))
            .apply();
    }
}
