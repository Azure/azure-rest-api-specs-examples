import com.azure.resourcemanager.peering.models.LegacyPeeringsKind;

/** Samples for LegacyPeerings List. */
public final class Main {
    /*
     * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-01-01/examples/ListLegacyPeerings.json
     */
    /**
     * Sample code: List legacy peerings.
     *
     * @param manager Entry point to PeeringManager.
     */
    public static void listLegacyPeerings(com.azure.resourcemanager.peering.PeeringManager manager) {
        manager
            .legacyPeerings()
            .list("peeringLocation0", LegacyPeeringsKind.EXCHANGE, null, com.azure.core.util.Context.NONE);
    }
}
