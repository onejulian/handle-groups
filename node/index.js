export const handler = async(event) => {
    const handleGroups = (passengersGroups) => {
        var passengersInt = passengersGroups.split(',').map(Number);
        var totalPassengers = passengersInt.reduce(function(a, b) {
            return a + b;
        });
        
        var answer = [];

        for (var i = 0; i < totalPassengers; i++) {
            var canCarryEveryone = true;
            var busCapacity = i+1;
            for (var j = 0; j < passengersInt.length; j++) {
                let pasajerosActuales = passengersInt[j];
            if (busCapacity < pasajerosActuales) {
                canCarryEveryone = false;
                break;
            }
            if (!checkGroups(fillGroups(passengersGroups), busCapacity) || (totalPassengers % busCapacity != 0)) {
                canCarryEveryone = false;
                break;
            }
            }
            if (canCarryEveryone) {
                answer.push(busCapacity);
            }
        }
        return answer.toString().replace("[", "").replace("]", "");
    };
    const fillGroups = (passengersGroups) => {
        let passengersInt = [];
        let passengersStr = passengersGroups.split(",");
        for (let i = 0; i < passengersStr.length; i++) {
            passengersInt.push(parseInt(passengersStr[i]));
        }
        return passengersInt;
    };
    const checkGroups = (groups, busCapacity) => {
        for (let i = 0; i < groups.length; i++) {
            let grupoActual = groups[i];
            if (grupoActual > busCapacity) {
                return false;
            }
            if (grupoActual == busCapacity) {
                groups.splice(i, 1);
                i--;
            }
            else if (grupoActual < busCapacity) {
                let passengers = grupoActual;
                groups.splice(i, 1);
                i--;
                while (passengers < busCapacity) {
                    if (i+1 < groups.length) {
                        passengers += groups[i+1];
                        groups.splice(i+1, 1);
                    }
                    else {
                        break;
                    }
                }
                if (passengers > busCapacity) {
                    return false;
                }
            }
        }
        return (groups.length == 0);
    };
    const response = {
        sizes: handleGroups(event.groups),
    };
    return response;
};
