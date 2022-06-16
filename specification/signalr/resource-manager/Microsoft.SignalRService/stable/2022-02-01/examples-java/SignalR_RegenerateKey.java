import com.azure.core.util.Context;
import com.azure.resourcemanager.signalr.models.KeyType;
import com.azure.resourcemanager.signalr.models.RegenerateKeyParameters;

/** Samples for SignalR RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalR_RegenerateKey.json
     */
    /**
     * Sample code: SignalR_RegenerateKey.
     *
     * @param manager Entry point to SignalRManager.
     */
    public static void signalRRegenerateKey(com.azure.resourcemanager.signalr.SignalRManager manager) {
        manager
            .signalRs()
            .regenerateKey(
                "myResourceGroup",
                "mySignalRService",
                new RegenerateKeyParameters().withKeyType(KeyType.PRIMARY),
                Context.NONE);
    }
}
