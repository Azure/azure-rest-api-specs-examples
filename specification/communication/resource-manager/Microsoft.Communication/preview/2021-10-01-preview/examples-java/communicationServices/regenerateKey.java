import com.azure.core.util.Context;
import com.azure.resourcemanager.communication.models.KeyType;
import com.azure.resourcemanager.communication.models.RegenerateKeyParameters;

/** Samples for CommunicationServices RegenerateKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2021-10-01-preview/examples/communicationServices/regenerateKey.json
     */
    /**
     * Sample code: Regenerate key.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void regenerateKey(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .communicationServices()
            .regenerateKey(
                "MyResourceGroup",
                "MyCommunicationResource",
                new RegenerateKeyParameters().withKeyType(KeyType.PRIMARY),
                Context.NONE);
    }
}
