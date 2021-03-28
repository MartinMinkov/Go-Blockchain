if [ -z $1 ]
then
    echo "Usage: ./set-peer.sh [set | unset]"
fi

if [[ $1 = 'set' ]]
then
    GENERATE_PEER_PORT='true'
else
    GENERATE_PEER_PORT=''
fi