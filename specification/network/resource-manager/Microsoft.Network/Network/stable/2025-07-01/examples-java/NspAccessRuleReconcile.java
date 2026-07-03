
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import java.io.IOException;

/**
 * Samples for NetworkSecurityPerimeterAccessRules Reconcile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/NspAccessRuleReconcile.json
     */
    /**
     * Sample code: NspAccessRuleReconcile.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void nspAccessRuleReconcile(com.azure.resourcemanager.network.NetworkManager manager)
        throws IOException {
        manager.serviceClient().getNetworkSecurityPerimeterAccessRules().reconcileWithResponse(
            "rg1", "nsp1", "profile1", "accessRuleName1", SerializerFactory.createDefaultManagementSerializerAdapter()
                .deserialize("{\"properties\":{}}", Object.class, SerializerEncoding.JSON),
            com.azure.core.util.Context.NONE);
    }
}
